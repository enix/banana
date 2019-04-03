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
        <h1>isSetup: {this.props.isSetup ? "OK" : "KO"} </h1>
        <button type="button" className="btn btn-primary">Primary</button>
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
