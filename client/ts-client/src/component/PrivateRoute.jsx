import React, { Component } from 'react';
import { Route, Redirect } from 'react-router-dom';
import { connect } from 'react-redux';

const PrivateRoute = props => {
  const { loginState } = props;
  return (
    <Route
      render={() =>
        loginState ? (
          <Component {...props} />
        ) : (
          <Redirect to={{ pathname: '/login' }} />
        )
      }
    />
  );
};

const mapStateToProps = state => {
  return {
    loginState: state.login.loginState
  };
};

export default connect(mapStateToProps)(PrivateRoute);
