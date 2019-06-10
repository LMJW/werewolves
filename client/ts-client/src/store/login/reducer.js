import actionTypes from '../constants';
export const loginReducer = (
	state = {
		loginState: false,
		username: '',
		session: '',
		loading: false,
		error: null
	},
	action
) => {
	switch (action.type) {
		case actionTypes.USER_LOGIN_START:
			return { ...state, loading: true, error: null };
		case actionTypes.USER_LOGIN_DONE:
			const { username } = action;
			return { ...state, loading: false, error: null, loginState: true, username };
		case actionTypes.USER_LOGIN_FAIL:
			const { error } = action;
			return { ...state, loginState: false, loading: false, error };
		case actionTypes.UPDATE_LOGIN_FORM_CONTEXT: {
			const { username } = action;
			return { ...state, username };
		}
		default:
			return state;
	}
};
