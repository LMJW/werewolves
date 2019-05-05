import {
  UserLoginAction,
  UserLoginState,
  USER_LOGIN,
  VALIDATE_SESSION
} from './types';

const initial_state: UserLoginState = {
  loginState: false,
  username: '',
  session: ''
};

export function login(
  state: UserLoginState = initial_state,
  action: UserLoginAction
): UserLoginState {
  switch (action.type) {
    case USER_LOGIN:
      return Object.assign({}, state, action.payload, { loginState: true });
    default:
      return state;
  }
}
