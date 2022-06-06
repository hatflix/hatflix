import { gql } from "apollo-boost";

export const GET_DISHES = gql`
     query all_dishes {
         dish {
             id
             name
             price
             cookTime
        }
    }
`;

export const GET_DISH_BY_ID = gql`
    query dish_by_id ($input : Int){
        dish(id: $input){
            id
            name
            price
            cookTime
            category {
                id
                name
            }
            restaurant {
                id
                name
                description
                address
            }
        }
    }
`;
