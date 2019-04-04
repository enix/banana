import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'

import ActionCreators from '../state/actions';
import { generateRestoreCmd } from '../helpers';

class App extends Component {

  componentDidMount() {
    this.props.actions.setupApp();
    console.log(generateRestoreCmd({
      name: 'etc',
      target: '/restored-etc',
      time: '20190404T130959Z',
      only: 'ok',
    }));
  }

  test = () => {
    this.props.actions.listBackupContainers();
    setTimeout(() => this.props.actions.listBackupsInContainer('banana-test2'), 1000);
  };

  render() {
    return (
      <div className="App">
        <h1>isSetup: {this.props.isSetup ? "OK" : "KO"} </h1>
        <button
          type="button"
          className="btn btn-primary"
          onClick={this.test}
        >
          Primary
        </button>
        <br />
        {JSON.stringify(this.props.containers, null, 2)}
      </div>
    );
  }
}

const mapStateToProps = state => ({
  isSetup: state.app.isSetup,
  containers: state.containers,
});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
