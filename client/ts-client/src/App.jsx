import React from 'react';
import './App.css';
import Login from './component/Login';
import Choice from './component/Choice';
import JoinGame from './component/JoinGame';
import Game from './component/Game';
import { Switch, Route } from 'react-router-dom';
import { connect } from 'react-redux';
import PrivateRoute from './component/PrivateRoute';

const App = props => {
  return (
    <div className='App'>
      <Switch>
        <Route path='/login' component={Login} />
        <PrivateRoute path='/' exact={true} component={Game} />
        <PrivateRoute path='/choice' component={Choice} />
        <PrivateRoute path='/join_game' component={JoinGame} />
      </Switch>
    </div>
  );
};

const mapStateToProps = state => ({
  loading: false,
  login: state.login,
  join: state.join
});

export default connect(mapStateToProps)(App);
