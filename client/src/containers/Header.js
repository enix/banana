import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'
import { Layout, Menu, Icon } from 'antd';

import Loading from '../components/Loading';
import ActionCreators from '../state/actions';

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
          <Menu.Item key='1'>Dashboard</Menu.Item>
        </Menu>
      </AntHeader>
    );
  }
}

const mapStateToProps = state => ({
  user: state.user,
});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Header);