import React, { Component } from 'react';
import { Icon } from 'antd';

class Loading extends Component {
  
  render() {
    return (
      <Icon type="loading-3-quarters" spin />
    );
  }
}

export default Loading;
