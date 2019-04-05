import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'

import Link from '../components/Link';
import List from '../components/List';
import Carret from '../components/Carret';
import ActionCreators from '../state/actions';

class Containers extends Component {

  renderItem = (item) => (
    <Link to={'/node/' + item.name}>
      <Carret />
      {item.name}
    </Link>
  )

  componentDidMount() {
    this.props.actions.listBackupContainers();
  }

  render() {
    return (
      <div className="Containers">
        <h2>Available nodes</h2>
        <List
          data={this.props.containers}
          renderItem={this.renderItem}
        />
      </div>
    );
  }
}

const mapStateToProps = state => ({
  containers: state.containers,
});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Containers);
