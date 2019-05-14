import { combineReducers } from 'redux';
import { loginReducer } from './login/reducer';
import { joinReducer } from './joinGame/reducer';

export const RootReducer = combineReducers({
  login: loginReducer,
  join: joinReducer
});
