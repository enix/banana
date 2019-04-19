import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'
import { Button } from 'antd';

import ActionCreators from '../state/actions';

class Header extends Component {

  render() {
    return (
      <Button>ok</Button>
      // <Layout>
      //   <Header style={{ position: 'fixed', zIndex: 1, width: '100%' }}>
      //     <div className="logo" />
      //     <Menu
      //       theme="dark"
      //       mode="horizontal"
      //       defaultSelectedKeys={['2']}
      //       style={{ lineHeight: '64px' }}
      //     >
      //       <Menu.Item key="1">nav 1</Menu.Item>
      //       <Menu.Item key="2">nav 2</Menu.Item>
      //       <Menu.Item key="3">nav 3</Menu.Item>
      //     </Menu>
      //   </Header>
      //   <Content style={{ padding: '0 50px', marginTop: 64 }}>
      //     <Breadcrumb style={{ margin: '16px 0' }}>
      //       <Breadcrumb.Item>Home</Breadcrumb.Item>
      //       <Breadcrumb.Item>List</Breadcrumb.Item>
      //       <Breadcrumb.Item>App</Breadcrumb.Item>
      //     </Breadcrumb>
      //     <div style={{ background: '#fff', padding: 24, minHeight: 380 }}>Content</div>
      //   </Content>
      //   <Footer style={{ textAlign: 'center' }}>
      //     Ant Design ©2018 Created by Ant UED
      //   </Footer>
      // </Layout>
    );
  }
}

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Header);
