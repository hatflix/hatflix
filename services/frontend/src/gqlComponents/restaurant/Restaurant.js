import React from "react";

const Restaurants = ({ loading, error, data }) => {
  if (loading) return "Loading...";
  if (error) return `Error! ${error.message}`;
  return (
    <ul>
      {data.restaurant.map(restaurant => (
        <a onClick={ () => {window.location.pathname === "/restaurant/" ?  window.location =  restaurant.id : window.location = "restaurant/"+restaurant.id}
        }>{restaurant.name}</a>
      ))} 
    </ul>
  );
};

const Restaurant = ({ loading, error, data }) => {
  if (loading) return "Loading...";
  if (error) return `Error! ${error.message}`;
  return (
    <>
      <h1>{data.restaurant[0].name}</h1>
      <h2>{data.restaurant[0].description}</h2>
      <h3>Horario de funcionamento: {data.restaurant[0].openHour} - {data.restaurant[0].closeHour}</h3>
      <h3>Dias de funcionamento: {data.restaurant[0].openDays.map(dia => dia+" ")}</h3>
      <h3>Telefone: {data.restaurant[0].phoneNumber}</h3>
      <h3>Endereço: {data.restaurant[0].address}</h3>
      <ul>
      {data.restaurant[0].dishes && data.restaurant[0].dishes.map(dish => (
        <a onClick={() => {
            window.location.pathname = "dish/"+dish.id
        }
        }>{dish.name} - ${dish.price}</a>
      ))}
    </ul>
    </>
  );
};

export { Restaurants, Restaurant };
