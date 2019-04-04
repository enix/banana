/* eslint no-script-url: "off" */

import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'

import ActionCreators from '../state/actions';

class List extends Component {

  componentDidMount() {
  }

  render() {
    return (
      <ul className="list-unstyled u-list">
        {this.props.data && this.props.data.map((item, index) => (
          <li key={index} onClick={() => this.props.onClick && this.props.onClick(item)}>
            {this.props.renderItem(item)}
          </li>
        ))}
      </ul>
    );
  }

  static propTypes = {
    data: PropTypes.array.isRequired,
    renderItem: PropTypes.func.isRequired,
    onClick: PropTypes.func,
  }
}

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(List);
