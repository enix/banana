import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'

import Link from '../components/Link';
import List from '../components/List';
import Carret from '../components/Carret';
import Loading from '../components/Loading';
import ActionCreators from '../state/actions';

class Agent extends Component {

  componentDidMount() {
    const { org, cn } = this.props.match.params;
    this.props.actions.getAgent(org, cn);
  }

  render() {
    if (!this.props.agent) {
      return <Loading />;
    }

    console.log(this.props.agent)

    return (
      <div className="Agent">
        <h2>Agent {this.props.agent.cn} from {this.props.agent.organization}</h2>
      </div>
    );
  }
}

const mapStateToProps = (state, props) => {
  const { org, cn } = props.match.params;

  if (!state.agents || !state.agents[org]) {
    return {};
  }

  return {
    agent: state.agents[org][cn]
  };
};

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Agent);
