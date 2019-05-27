import React, { Component } from 'react';
import { connect } from 'react-redux'
import { mapDispatchToProps } from 'redux-saga-wrapper';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { Layout } from 'antd';

import Settings from '../containers/Settings';
import Agents from '../containers/Agents';
import Agent from '../containers/Agent';
import Header from './Header';

const { Content } = Layout;

class App extends Component {

  componentDidMount() {
    this.props.actions.setupApp();
  }

  render() {
    return (
      <Router>
        <div className="App">
          <Layout className="layout">
            <Header />
            <Content style={{ padding: '40px', background: '#fff' }}>
              <Route exact path='/' component={Agents} />
              <Route exact path='/settings' component={Settings} />
              <Route exact path='/agent/:org/:cn' component={Agent} />
            </Content>
          </Layout>
        </div>
      </Router>
    );
  }
}

const mapStateToProps = state => ({});

export default connect(mapStateToProps, mapDispatchToProps)(App);
