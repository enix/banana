import React, { Component } from 'react';
import PropTypes from 'prop-types';

class List extends Component {

  render() {
    return (
      <ul className="list-unstyled u-list">
        {this.props.data && this.props.data.map((item, index) => (
          <li key={index} onClick={() => this.props.onClick && this.props.onClick(item)}>
            {this.props.renderItem(item, index)}
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

export default List;
