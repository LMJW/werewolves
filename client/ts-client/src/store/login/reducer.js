import { USER_LOGIN } from './types';

export const loginReducer = (
  state = {
    loginState: false,
    username: '',
    session: ''
  },
  action
) => {
  switch (action.type) {
    case USER_LOGIN:
      return Object.assign({}, state, action.payload, { loginState: true });
    default:
      return state;
  }
};
