import React, { Component } from 'react';
import PropsType from 'prop-types';
import { Table } from 'antd';

import './JsonTable.less';

class JsonTable extends Component {

  columns = [
    {
      title: 'Key',
      dataIndex: 'key',
      key: 'key',
      render: value => (value[0].toUpperCase() + value.slice(1)).replace(/_/g, ' '),
    },
    {
      title: 'Value',
      dataIndex: 'value',
      key: 'value',
      render: value => (
        <pre className='JsonTable__value'>
          {typeof value === 'object' ? JSON.stringify(value, null, 2) : value}
        </pre>
      ),
    },
  ]

  getDataSource = () => {
    console.log(Object.keys(this.props.data).map(key => ({
      key,
      value: this.props.data[key],
    })))
    return Object.keys(this.props.data).map(key => ({
      key,
      value: this.props.data[key],
    }));
  }

  render() {
    return (
      <Table
        bordered
        columns={this.columns}
        dataSource={this.getDataSource()}
        pagination={false}
      />
    );
  }

  static propsType = {
    data: PropsType.object.isRequired,
  }
}

export default JsonTable;
