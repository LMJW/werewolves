export interface UserInfo {
  username: string;
  session: string;
}

export interface UserLoginState {
  loginState: boolean;
  username: string;
  session: string;
}

export const USER_LOGIN = 'USER_LOGIN';
export const VALIDATE_SESSION = 'VALIDATE_SESSION';

export interface UserLoginAction {
  type: typeof USER_LOGIN;
  payload: UserInfo;
}
