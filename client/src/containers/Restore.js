import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'

import ActionCreators from '../state/actions';
import { generateRestoreCmd } from '../helpers';

class Restore extends Component {

  state = {
    cmd: '',
  }

  async componentDidMount() {
    const { treeName, time } = this.props.match.params;
    const cmd = generateRestoreCmd({
      name: treeName,
      time,
      target: '/path/to/restore',
    });

    this.setState({ cmd });
  }

  render() {
    return (
      <div className="Restore">
        {this.state.cmd}
      </div>
    );
  }
}

const mapStateToProps = (state, props) => ({});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Restore);
