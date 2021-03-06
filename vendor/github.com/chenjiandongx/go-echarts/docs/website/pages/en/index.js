/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

const React = require('react');

const CompLibrary = require('../../core/CompLibrary.js');

const Container = CompLibrary.Container;
const GridBlock = CompLibrary.GridBlock;

class HomeSplash extends React.Component {
  render() {
    const {siteConfig, language = ''} = this.props;
    const {baseUrl, docsUrl} = siteConfig;
    const docsPart = `${docsUrl ? `${docsUrl}/` : ''}`;
    const langPart = `${language ? `${language}/` : ''}`;
    const docUrl = doc => `${baseUrl}${docsPart}${langPart}${doc}`;

    const SplashContainer = props => (
      <div className="homeContainer">
        <div className="homeSplashFade">
          <div className="wrapper homeWrapper">{props.children}</div>
        </div>
      </div>
    );

    const Logo = props => (
      <div className="projectLogo">
        <img src={props.img_src} alt="Project Logo"/>
      </div>
    );

    const ProjectTitle = () => (
      <h2 className="projectTitle">
        {siteConfig.title}
        <small>{siteConfig.tagline}</small>
      </h2>
    );

    const PromoSection = props => (
      <div className="section promoSection">
        <div className="promoRow">
          <div className="pluginRowBlock">{props.children}</div>
        </div>
      </div>
    );

    const Button = props => (
      <div className="pluginWrapper buttonWrapper">
        <a className="button" href={props.href} target={props.target}>
          {props.children}
        </a>
      </div>
    );

    return (
      <SplashContainer>
        <Logo img_src={`${baseUrl}img/logo.png`}/>
        <div className="inner">
          <ProjectTitle siteConfig={siteConfig} />
          <PromoSection>
            <Button href="docs/overview">Get start</Button>
          </PromoSection>
        </div>
      </SplashContainer>
    );
  }
}

class Index extends React.Component {
  render() {
    const {config: siteConfig, language = ''} = this.props;
    const {baseUrl} = siteConfig;

    const Block = props => (
      <Container
        padding={['bottom', 'top']}
        id={props.id}
        background={props.background}>
        <GridBlock
          align="center"
          contents={props.children}
          layout={props.layout}
        />
      </Container>
    );

    const TryOut = () => (
      <Block id="try" background="light">
        {[
          {
            content: '?????? HTTP ?????????????????????????????????',
            image: `${baseUrl}img/run_example.png`,
            imageAlign: 'left',
            title: '5 ????????????',
          },
        ]}
      </Block>
    );

    const LearnHow = () => (
      <Block>
        {[
          {
            content: '????????? 5 ?????????????????????????????????????????????',
            image: `${baseUrl}img/quickstart.png`,
            imageAlign: 'right',
            title: '1 ????????????',
          },
        ]}
      </Block>
    );

    const Features = () => (
      <Block layout="fourColumn" background="light">
        {[
          {
            content: '????????????????????? ????????????????????????',
            image: `${baseUrl}img/concise.png`,
            imageAlign: 'top',
            title: '??????',
          },
          {
            content: '????????? 25+ ??????????????? ????????????',
            image: `${baseUrl}img/abundant.png`,
            imageAlign: 'top',
            title: '??????',
          },
          {
            content: '???????????????????????? ?????????????????????????????????',
            image: `${baseUrl}img/flexible.png`,
            imageAlign: 'top',
            title: '??????',
          },
        ]}
      </Block>
    );

    return (
      <div>
        <HomeSplash siteConfig={siteConfig} language={language} />
        <div className="mainContainer">
          <Features />
          <LearnHow />
          <TryOut />
        </div>
      </div>
    );
  }
}

module.exports = Index;
