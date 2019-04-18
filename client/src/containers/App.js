import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';

import Agents from '../containers/Agents';
import Agent from '../containers/Agent';
import ActionCreators from '../state/actions';

class App extends Component {

  componentDidMount() {
    this.props.actions.setupApp();
  }

  render() {
    return (
      <Router>
        <div className="App" style={{ padding: 20 }}>
          <Link to='/'>
            <h1>Banana</h1>
          </Link>
          <Route exact path='/' component={Agents} />
          <Route exact path='/agent/:org/:cn' component={Agent} />
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
