import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'
import { Link as LowLevelLink } from 'react-router-dom';

import ActionCreators from '../state/actions';

class Link extends Component {

  render() {
    return (
      <LowLevelLink
        to={this.props.to}
        style={{ textDecoration: 'none', color: 'inherit' }}
        className="u-list__link"
        href='javascript:void(0)'
      >
        {this.props.children}
      </LowLevelLink>
    );
  }

  static propTypes = {
    to: PropTypes.string.isRequired,
  }
}

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Link);
