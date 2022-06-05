import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { Restaurants, Restaurant } from "./Restaurant";
import { Link } from 'react-router-dom'

import { 
  GET_RESTAURANTES, 
  GET_RESTAURANTES_BY_ID
} from "./Query";

const GetAllRestaurants = () => {
  const { loading, error, data, refetch } = useQuery(GET_RESTAURANTES, {
    fetchPolicy: "cache-and-network",
  });

  return (
    <>
      {" "}
      <Restaurants error={error} loading={loading} data={data} />
      <Link to="/">Home</Link>
    </>
  );
};

const GetRestaurantsById = ({filter}) => {
  const { loading, error, data, refetch } = useQuery(GET_RESTAURANTES_BY_ID, {
    fetchPolicy: "cache-and-network",
    variables: { input: filter },
  });
  return (
    <>
      {" "}
      <Restaurant error={error} loading={loading} data={data} />
      <Link to="/">Home</Link>
    </>
  );
};

export { GetAllRestaurants, GetRestaurantsById };
