import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { Categories, Category } from "./Category";
import { Link } from 'react-router-dom'

import {
    GET_CATEGORIES,
    GET_CATEGORY_BY_ID
} from "./Querys";

const GetAllCategories = () => {
    const { loading, error, data, refetch } = useQuery(GET_CATEGORIES, {
        fetchPolicy: "cache-and-network",
    });

    return (
        <>
            {" "}
            <Categories error={error} loading={loading} data={data} />
            <Link to="/">Home</Link>
        </>
    );
};

const GetCategoryByID = ({filter}) => {
    const { loading, error, data, refetch } = useQuery(GET_CATEGORY_BY_ID, {
        fetchPolicy: "cache-and-network",
        variables: { input: filter },
    });
    return (
        <>
            {" "}
            <Category error={error} loading={loading} data={data} />
            <Link to="/">Home</Link>
        </>
    );
};

export { GetAllCategories, GetCategoryByID };
