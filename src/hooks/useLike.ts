import { useDispatch , useSelector } from 'react-redux';
import { bindActionCreators } from '@zealforchange/conciseredux'
import { like, likeStateTypes  } from 'src/reducer/like';

export function useLike() {

const dispatchWithLike = bindActionCreators(
like.getCallBackAll(),useDispatch())

const curStateWithLikeForRedux = like.getCurState()

const curStateWithLike = useSelector((state: {like:likeStateTypes }) => {
return state.like
})

return {dispatchWithLike , curStateWithLike , curStateWithLikeForRedux}

}