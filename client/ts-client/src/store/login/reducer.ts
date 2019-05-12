import {
  UserLoginAction,
  UserLoginState,
  USER_LOGIN,
  VALIDATE_SESSION
} from './types';

import { Reducer } from 'redux';

const initial_state: UserLoginState = {
  loginState: false,
  username: '',
  session: ''
};

export const loginReducer: Reducer<UserLoginState> = (
  state = initial_state,
  action
): UserLoginState => {
  switch (action.type) {
    case USER_LOGIN:
      return Object.assign({}, state, action.payload, { loginState: true });
    default:
      return state;
  }
};
