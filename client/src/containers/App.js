import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { Layout } from 'antd';

import Agents from '../containers/Agents';
import Agent from '../containers/Agent';
import Header from './Header';
import ActionCreators from '../state/actions';

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
              <Route exact path='/agent/:org/:cn' component={Agent} />
            </Content>
          </Layout>
        </div>
      </Router>
    );
  }
}

const mapStateToProps = state => ({
  isSetup: state.app.isSetup,
});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
