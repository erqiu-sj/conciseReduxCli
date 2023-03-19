import { useDispatch , useSelector } from 'react-redux';
import { bindActionCreators } from '@zealforchange/conciseredux'
import { action, actionStateTypes  } from 'src/reducer/action';

export function useAction() {

const dispatchWithAction = bindActionCreators(
action.getCallBackAll(),useDispatch())

const curStateWithActionForRedux = action.getCurState()

const curStateWithAction = useSelector((state: {action:actionStateTypes }) => {
return state.action
})

return {dispatchWithAction , curStateWithAction , curStateWithActionForRedux}

}