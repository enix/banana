import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'

import ActionCreators from '../state/actions';

class App extends Component {
  componentDidMount() {
    this.props.actions.setupApp();
  }

  render() {
    return (
      <div className="App">
        <span>isSetup: {this.props.isSetup ? "OK" : "KO"} </span>
      </div>
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
