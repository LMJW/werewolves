import React from 'react';
import './App.css';
import Login from './component/Login';
import Choice from './component/Choice';
import JoinGame from './component/JoinGame';
import Game from './component/Game';
import { Switch, Route } from 'react-router-dom';
import { AppState } from './store/index';
import { connect } from 'react-redux';

interface AppProps {
  loading: boolean;
  error?: string;
}

type Props = AppProps & AppState;

const App = (props: Props) => {
  return (
    <div className='App'>
      <Switch>
        <Route path='/login' component={Login} />
        <Route path='/' exact={true} component={Game} />
        <Route path='/choice' component={Choice} />
        <Route path='/join_game' component={JoinGame} />
      </Switch>
    </div>
  );
};

const mapStateToProps = (state: AppState): Props => ({
  loading: false,
  login: state.login,
  join: state.join
});

export default connect(mapStateToProps)(App);
