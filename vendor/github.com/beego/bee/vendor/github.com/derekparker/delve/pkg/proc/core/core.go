package core

import (
	"errors"
	"fmt"
	"go/ast"
	"io"
	"sync"

	"github.com/derekparker/delve/pkg/proc"
)

// A SplicedMemory represents a memory space formed from multiple regions,
// each of which may override previously regions. For example, in the following
// core, the program text was loaded at 0x400000:
// Start               End                 Page Offset
// 0x0000000000400000  0x000000000044f000  0x0000000000000000
// but then it's partially overwritten with an RW mapping whose data is stored
// in the core file:
// Type           Offset             VirtAddr           PhysAddr
//                FileSiz            MemSiz              Flags  Align
// LOAD           0x0000000000004000 0x000000000049a000 0x0000000000000000
//                0x0000000000002000 0x0000000000002000  RW     1000
// This can be represented in a SplicedMemory by adding the original region,
// then putting the RW mapping on top of it.
type SplicedMemory struct {
	readers []readerEntry
}

type readerEntry struct {
	offset uintptr
	length uintptr
	reader proc.MemoryReader
}

// Add adds a new region to the SplicedMemory, which may override existing regions.
func (r *SplicedMemory) Add(reader proc.MemoryReader, off, length uintptr) {
	if length == 0 {
		return
	}
	end := off + length - 1
	newReaders := make([]readerEntry, 0, len(r.readers))
	add := func(e readerEntry) {
		if e.length == 0 {
			return
		}
		newReaders = append(newReaders, e)
	}
	inserted := false
	// Walk through the list of regions, fixing up any that overlap and inserting the new one.
	for _, entry := range r.readers {
		entryEnd := entry.offset + entry.length - 1
		switch {
		case entryEnd < off:
			// Entry is completely before the new region.
			add(entry)
		case end < entry.offset:
			// Entry is completely after the new region.
			if !inserted {
				add(readerEntry{off, length, reader})
				inserted = true
			}
			add(entry)
		case off <= entry.offset && entryEnd <= end:
			// Entry is completely overwritten by the new region. Drop.
		case entry.offset < off && entryEnd <= end:
			// New region overwrites the end of the entry.
			entry.length = off - entry.offset
			add(entry)
		case off <= entry.offset && end < entryEnd:
			// New reader overwrites the beginning of the entry.
			if !inserted {
				add(readerEntry{off, length, reader})
				inserted = true
			}
			overlap := entry.offset - off
			entry.offset += overlap
			entry.length -= overlap
			add(entry)
		case entry.offset < off && end < entryEnd:
			// New region punches a hole in the entry. Split it in two and put the new region in the middle.
			add(readerEntry{entry.offset, off - entry.offset, entry.reader})
			add(readerEntry{off, length, reader})
			add(readerEntry{end + 1, entryEnd - end, entry.reader})
			inserted = true
		default:
			panic(fmt.Sprintf("Unhandled case: existing entry is %v len %v, new is %v len %v", entry.offset, entry.length, off, length))
		}
	}
	if !inserted {
		newReaders = append(newReaders, readerEntry{off, length, reader})
	}
	r.readers = newReaders
}

// ReadMemory implements MemoryReader.ReadMemory.
func (r *SplicedMemory) ReadMemory(buf []byte, addr uintptr) (n int, err error) {
	started := false
	for _, entry := range r.readers {
		if entry.offset+entry.length < addr {
			if !started {
				continue
			}
			return n, fmt.Errorf("hit unmapped area at %v after %v bytes", addr, n)
		}

		// Don't go past the region.
		pb := buf
		if addr+uintptr(len(buf)) > entry.offset+entry.length {
			pb = pb[:entry.offset+entry.length-addr]
		}
		pn, err := entry.reader.ReadMemory(pb, addr)
		n += pn
		if err != nil || pn != len(pb) {
			return n, err
		}
		buf = buf[pn:]
		addr += uintptr(pn)
		if len(buf) == 0 {
			// Done, don't bother scanning the rest.
			return n, nil
		}
	}
	if n == 0 {
		return 0, fmt.Errorf("offset %v did not match any regions", addr)
	}
	return n, nil
}

// OffsetReaderAt wraps a ReaderAt into a MemoryReader, subtracting a fixed
// offset from the address. This is useful to represent a mapping in an address
// space. For example, if program text is mapped in at 0x400000, an
// OffsetReaderAt with offset 0x400000 can be wrapped around file.Open(program)
// to return the results of a read in that part of the address space.
type OffsetReaderAt struct {
	reader io.ReaderAt
	offset uintptr
}

// ReadMemory will read the memory at addr-offset.
func (r *OffsetReaderAt) ReadMemory(buf []byte, addr uintptr) (n int, err error) {
	return r.reader.ReadAt(buf, int64(addr-r.offset))
}

// Process represents a core file.
type Process struct {
	bi                *proc.BinaryInfo
	core              *Core
	breakpoints       proc.BreakpointMap
	currentThread     *Thread
	selectedGoroutine *proc.G
	common            proc.CommonProcess
}

// Thread represents a thread in the core file being debugged.
type Thread struct {
	th     *LinuxPrStatus
	fpregs []proc.Register
	p      *Process
	common proc.CommonThread
}

var (
	// ErrWriteCore is returned when attempting to write to the core
	// process memory.
	ErrWriteCore = errors.New("can not write to core process")

	// ErrShortRead is returned on a short read.
	ErrShortRead = errors.New("short read")

	// ErrContinueCore is returned when trying to continue execution of a core process.
	ErrContinueCore = errors.New("can not continue execution of core process")

	// ErrChangeRegisterCore is returned when trying to change register values for core files.
	ErrChangeRegisterCore = errors.New("can not change register values of core process")
)

// OpenCore will open the core file and return a Process struct.
func OpenCore(corePath, exePath string) (*Process, error) {
	core, err := readCore(corePath, exePath)
	if err != nil {
		return nil, err
	}
	p := &Process{
		core:        core,
		breakpoints: proc.NewBreakpointMap(),
		bi:          proc.NewBinaryInfo("linux", "amd64"),
	}
	for _, thread := range core.Threads {
		thread.p = p
	}

	var wg sync.WaitGroup
	err = p.bi.LoadBinaryInfo(exePath, core.entryPoint, &wg)
	wg.Wait()
	if err == nil {
		err = p.bi.LoadError()
	}
	if err != nil {
		return nil, err
	}

	for _, th := range p.core.Threads {
		p.currentThread = th
		break
	}
	p.selectedGoroutine, _ = proc.GetG(p.CurrentThread())

	return p, nil
}

// BinInfo will return the binary info.
func (p *Process) BinInfo() *proc.BinaryInfo {
	return p.bi
}

// Recorded returns whether this is a live or recorded process. Always returns true for core files.
func (p *Process) Recorded() (bool, string) { return true, "" }

// Restart will only return an error for core files, as they are not executing.
func (p *Process) Restart(string) error { return ErrContinueCore }

// Direction will only return an error as you cannot continue a core process.
func (p *Process) Direction(proc.Direction) error { return ErrContinueCore }

// When does not apply to core files, it is to support the Mozilla 'rr' backend.
func (p *Process) When() (string, error) { return "", nil }

// Checkpoint for core files returns an error, there is no execution of a core file.
func (p *Process) Checkpoint(string) (int, error) { return -1, ErrContinueCore }

// Checkpoints returns nil on core files, you cannot set checkpoints when debugging core files.
func (p *Process) Checkpoints() ([]proc.Checkpoint, error) { return nil, nil }

// ClearCheckpoint clears a checkpoint, but will only return an error for core files.
func (p *Process) ClearCheckpoint(int) error { return errors.New("checkpoint not found") }

// ReadMemory will return memory from the core file at the specified location and put the
// read memory into `data`, returning the length read, and returning an error if
// the length read is shorter than the length of the `data` buffer.
func (t *Thread) ReadMemory(data []byte, addr uintptr) (n int, err error) {
	n, err = t.p.core.ReadMemory(data, addr)
	if err == nil && n != len(data) {
		err = ErrShortRead
	}
	return n, err
}

// WriteMemory will only return an error for core files, you cannot write
// to the memory of a core process.
func (t *Thread) WriteMemory(addr uintptr, data []byte) (int, error) {
	return 0, ErrWriteCore
}

// Location returns the location of this thread based on
// the value of the instruction pointer register.
func (t *Thread) Location() (*proc.Location, error) {
	f, l, fn := t.p.bi.PCToLine(t.th.Reg.Rip)
	return &proc.Location{PC: t.th.Reg.Rip, File: f, Line: l, Fn: fn}, nil
}

// Breakpoint returns the current breakpoint this thread is stopped at.
// For core files this always returns an empty BreakpointState struct, as
// there are no breakpoints when debugging core files.
func (t *Thread) Breakpoint() proc.BreakpointState {
	return proc.BreakpointState{}
}

// ThreadID returns the ID for this thread.
func (t *Thread) ThreadID() int {
	return int(t.th.Pid)
}

// Registers returns the current value of the registers for this thread.
func (t *Thread) Registers(floatingPoint bool) (proc.Registers, error) {
	r := &Registers{&t.th.Reg, nil}
	if floatingPoint {
		r.fpregs = t.fpregs
	}
	return r, nil
}

// RestoreRegisters will only return an error for core files,
// you cannot change register values for core files.
func (t *Thread) RestoreRegisters(proc.Registers) error {
	return ErrChangeRegisterCore
}

// Arch returns the architecture the target is built for and executing on.
func (t *Thread) Arch() proc.Arch {
	return t.p.bi.Arch
}

// BinInfo returns information about the binary.
func (t *Thread) BinInfo() *proc.BinaryInfo {
	return t.p.bi
}

// StepInstruction will only return an error for core files,
// you cannot execute a core file.
func (t *Thread) StepInstruction() error {
	return ErrContinueCore
}

// Blocked will return false always for core files as there is
// no execution.
func (t *Thread) Blocked() bool {
	return false
}

// SetCurrentBreakpoint will always just return nil
// for core files, as there are no breakpoints in core files.
func (t *Thread) SetCurrentBreakpoint() error {
	return nil
}

// Common returns a struct containing common information
// across thread implementations.
func (t *Thread) Common() *proc.CommonThread {
	return &t.common
}

// SetPC will always return an error, you cannot
// change register values when debugging core files.
func (t *Thread) SetPC(uint64) error {
	return ErrChangeRegisterCore
}

// SetSP will always return an error, you cannot
// change register values when debugging core files.
func (t *Thread) SetSP(uint64) error {
	return ErrChangeRegisterCore
}

// SetDX will always return an error, you cannot
// change register values when debugging core files.
func (t *Thread) SetDX(uint64) error {
	return ErrChangeRegisterCore
}

// Breakpoints will return all breakpoints for the process.
func (p *Process) Breakpoints() *proc.BreakpointMap {
	return &p.breakpoints
}

// ClearBreakpoint will always return an error as you cannot set or clear
// breakpoints on core files.
func (p *Process) ClearBreakpoint(addr uint64) (*proc.Breakpoint, error) {
	return nil, proc.NoBreakpointError{Addr: addr}
}

// ClearInternalBreakpoints will always return nil and have no
// effect since you cannot set breakpoints on core files.
func (p *Process) ClearInternalBreakpoints() error {
	return nil
}

// ContinueOnce will always return an error because you
// cannot control execution of a core file.
func (p *Process) ContinueOnce() (proc.Thread, error) {
	return nil, ErrContinueCore
}

// StepInstruction will always return an error
// as you cannot control execution of a core file.
func (p *Process) StepInstruction() error {
	return ErrContinueCore
}

// RequestManualStop will return nil and have no effect
// as you cannot control execution of a core file.
func (p *Process) RequestManualStop() error {
	return nil
}

// CheckAndClearManualStopRequest will always return false and
// have no effect since there are no manual stop requests as
// there is no controlling execution of a core file.
func (p *Process) CheckAndClearManualStopRequest() bool {
	return false
}

// CurrentThread returns the current active thread.
func (p *Process) CurrentThread() proc.Thread {
	return p.currentThread
}

// Detach will always return nil and have no
// effect as you cannot detach from a core file
// and have it continue execution or exit.
func (p *Process) Detach(bool) error {
	return nil
}

// Valid returns whether the process is active. Always returns true
// for core files as it cannot exit or be otherwise detached from.
func (p *Process) Valid() (bool, error) {
	return true, nil
}

// Common returns common information across Process
// implementations.
func (p *Process) Common() *proc.CommonProcess {
	return &p.common
}

// Pid returns the process ID of this process.
func (p *Process) Pid() int {
	return p.core.Pid
}

// ResumeNotify is a no-op on core files as we cannot
// control execution.
func (p *Process) ResumeNotify(chan<- struct{}) {
}

// SelectedGoroutine returns the current active and selected
// goroutine.
func (p *Process) SelectedGoroutine() *proc.G {
	return p.selectedGoroutine
}

// SetBreakpoint will always return an error for core files as you cannot write memory or control execution.
func (p *Process) SetBreakpoint(addr uint64, kind proc.BreakpointKind, cond ast.Expr) (*proc.Breakpoint, error) {
	return nil, ErrWriteCore
}

// SwitchGoroutine will change the selected and active goroutine.
func (p *Process) SwitchGoroutine(gid int) error {
	g, err := proc.FindGoroutine(p, gid)
	if err != nil {
		return err
	}
	if g == nil {
		// user specified -1 and selectedGoroutine is nil
		return nil
	}
	if g.Thread != nil {
		return p.SwitchThread(g.Thread.ThreadID())
	}
	p.selectedGoroutine = g
	return nil
}

// SwitchThread will change the selected and active thread.
func (p *Process) SwitchThread(tid int) error {
	if th, ok := p.core.Threads[tid]; ok {
		p.currentThread = th
		p.selectedGoroutine, _ = proc.GetG(p.CurrentThread())
		return nil
	}
	return fmt.Errorf("thread %d does not exist", tid)
}

// ThreadList will return a list of all threads currently in the process.
func (p *Process) ThreadList() []proc.Thread {
	r := make([]proc.Thread, 0, len(p.core.Threads))
	for _, v := range p.core.Threads {
		r = append(r, v)
	}
	return r
}

// FindThread will return the thread with the corresponding thread ID.
func (p *Process) FindThread(threadID int) (proc.Thread, bool) {
	t, ok := p.core.Threads[threadID]
	return t, ok
}

// Registers represents the CPU registers.
type Registers struct {
	*LinuxCoreRegisters
	fpregs []proc.Register
}

// Slice will return a slice containing all registers and their values.
func (r *Registers) Slice() []proc.Register {
	var regs = []struct {
		k string
		v uint64
	}{
		{"Rip", r.Rip},
		{"Rsp", r.Rsp},
		{"Rax", r.Rax},
		{"Rbx", r.Rbx},
		{"Rcx", r.Rcx},
		{"Rdx", r.Rdx},
		{"Rdi", r.Rdi},
		{"Rsi", r.Rsi},
		{"Rbp", r.Rbp},
		{"R8", r.R8},
		{"R9", r.R9},
		{"R10", r.R10},
		{"R11", r.R11},
		{"R12", r.R12},
		{"R13", r.R13},
		{"R14", r.R14},
		{"R15", r.R15},
		{"Orig_rax", r.Orig_rax},
		{"Cs", r.Cs},
		{"Eflags", r.Eflags},
		{"Ss", r.Ss},
		{"Fs_base", r.Fs_base},
		{"Gs_base", r.Gs_base},
		{"Ds", r.Ds},
		{"Es", r.Es},
		{"Fs", r.Fs},
		{"Gs", r.Gs},
	}
	out := make([]proc.Register, 0, len(regs))
	for _, reg := range regs {
		if reg.k == "Eflags" {
			out = proc.AppendEflagReg(out, reg.k, reg.v)
		} else {
			out = proc.AppendQwordReg(out, reg.k, reg.v)
		}
	}
	out = append(out, r.fpregs...)
	return out
}

// Copy will return a copy of the registers that is guarenteed
// not to change.
func (r *Registers) Copy() proc.Registers {
	return r
}
