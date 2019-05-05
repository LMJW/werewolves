import React from 'react';
import './App.css';
import { Provider } from 'react-redux';
import Login from './component/Login';
import { createStore, applyMiddleware } from 'redux';
import { RootReducer } from './store/index';
import { composeWithDevTools } from 'redux-devtools-extension';
import { createLogger } from 'redux-logger';

const store = createStore(
  RootReducer,
  composeWithDevTools(applyMiddleware(createLogger()))
);

const App: React.FC = () => {
  return (
    <Provider store={store}>
      <div className='App'>
        <Login />
      </div>
    </Provider>
  );
};

export default App;
