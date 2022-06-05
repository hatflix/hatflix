import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { Dishes, Dish } from "./Dish";
import { Link } from 'react-router-dom'

import {
    GET_DISHES,
    GET_DISH_BY_ID
} from "./Querys";

const GetAllDishes = () => {
    const { loading, error, data, refetch } = useQuery(GET_DISHES, {
        fetchPolicy: "cache-and-network",
    });

    return (
        <>
            {" "}
            <Dishes error={error} loading={loading} data={data} />
            <Link to="/">Home</Link>
        </>
    );
};

const GetDishById = ({filter}) => {
    const { loading, error, data, refetch } = useQuery(GET_DISH_BY_ID, {
        fetchPolicy: "cache-and-network",
        variables: { input: filter },
    });
    return (
        <>
            {" "}
            <Dish error={error} loading={loading} data={data} />
            <Link to="/">Home</Link>
        </>
    );
};

export { GetAllDishes, GetDishById };
