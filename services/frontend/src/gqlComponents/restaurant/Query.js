import { gql } from "apollo-boost";

export const GET_RESTAURANTES = gql`
    query teste{
        restaurant{
            id
            openDays
            name
            description
            phoneNumber
            address
        }
    }
`;

export const GET_RESTAURANTES_BY_ID = gql`
    query teste ($input : Int){
        restaurant(id: $input){
            id
            openHour
            closeHour
            openDays
            name
            description
            phoneNumber
            address
            dishes{
                id
                name
                price
            }
        }
    }
`;
