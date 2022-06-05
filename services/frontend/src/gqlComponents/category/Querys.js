import { gql } from "apollo-boost";

export const GET_CATEGORIES = gql`
    query all_categories{
        category{
            id
            name
        }
    }
`;

export const GET_CATEGORY_BY_ID = gql`
    query category_by_id ($input: Int) {
        category (id: $input) {
            id
            name
            restaurants {
                id
                name
                description
                address
            }
        }
    }
`;
