import { CreateReducer , getAllValsWithActionCollectionHelper } from '@zealforchange/conciseredux'
 

 
 export const  likeAction = {
 ACTIONS: 'actions',
 PATAS: 'patas'
}
 
 export type likeActionTypes =  getAllValsWithActionCollectionHelper<typeof likeAction>
 
 export type  likeStateTypes  = {}
 
 export type  likeActionPayloadTypes = {}
 
 export const like = new CreateReducer<likeStateTypes,likeActionPayloadTypes,likeActionTypes>({}) 
  .addAction('actions', (state, action) => {
    return { ...state }
  } .addAction('patas', (state, action) => {
    return { ...state }
  } 
 .setReducerKey('like')

