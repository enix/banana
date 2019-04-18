import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'

import Link from '../components/Link';
import List from '../components/List';
import Carret from '../components/Carret';
import Loading from '../components/Loading';
import ActionCreators from '../state/actions';

class Agents extends Component {

  renderItem = (item) => (
    <Link to={`/agent/${item.organization}/${item.cn}`}>
      <Carret />
      <b>{item.organization}</b> / {item.cn}
    </Link>
  )

  componentDidMount() {
    this.props.actions.listAgents();
  }

  render() {
    if (!this.props.agents) {
      return <Loading />
    }

    return (
      <div className="Agents">
        <h2>Available agents</h2>
        <List
          data={this.props.agents}
          renderItem={this.renderItem}
        />
      </div>
    );
  }
}

const mapStateToProps = state => ({
  agents: state.agentList,
});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Agents);
