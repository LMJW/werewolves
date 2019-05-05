import { combineReducers } from 'redux';
import { login } from './login/reducer';

export const RootReducer = combineReducers({ login });
