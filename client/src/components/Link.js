import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Link as LowLevelLink } from 'react-router-dom';

class Link extends Component {

  render() {
    return (
      <LowLevelLink className="u-list__link" {...this.props}>
        {this.props.children}
      </LowLevelLink>
    );
  }

  static propTypes = {
    to: PropTypes.string.isRequired,
  }
}

export default Link;
