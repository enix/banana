import React, { Component } from 'react';
import { connect } from 'react-redux';
import { mapDispatchToProps } from 'redux-saga-wrapper';
import { Layout, Menu, Icon } from 'antd';
import { Link } from 'react-router-dom';

import Loading from '../components/Loading';

import './Header.less';

const { Header: AntHeader } = Layout;

class Header extends Component {

  componentDidMount() {
    this.props.actions.pingApi();
  }

  render() {
    return (
      <AntHeader className='Header'>
        <img src='/img/logo.svg' className='logo' alt='logo' />
        <div className='loggedAs'>
          {!this.props.user.name ? <Loading /> : (
            <span>
              <Icon type="user" />
              <span> {this.props.user.organization} / {this.props.user.name}</span>
            </span>
          )}
        </div>
        <Menu
          theme='dark'
          mode='horizontal'
          defaultSelectedKeys={['1']}
          style={{ lineHeight: '64px' }}
        >
          <Menu.Item key='1'>
            <Link to='/'>Dashboard</Link>
          </Menu.Item>
          <Menu.Item key='2'>
            <Link to='/settings'>Settings</Link>
          </Menu.Item>
        </Menu>
      </AntHeader>
    );
  }
}

const mapStateToProps = state => ({
  user: state.user,
});

export default connect(mapStateToProps, mapDispatchToProps)(Header);
