import React from "react";

const Dishes = ({ loading, error, data }) => {
    if (loading) return "Loading...";
    if (error) return `Error! ${error.message}`;
    return (
        <ul>
            {data.dish && data.dish.map(dish => (
                <a onClick={ () => {window.location.pathname === "/dish/" ?  window.location =  dish.id : window.location = "dish/"+dish.id}
                }>{dish.name}</a>
            ))}
        </ul>
    );
};

const Dish = ({ loading, error, data }) => {
    if (loading) return "Loading...";
    if (error) return `Error! ${error.message}`;
    return (
        <>
            <h2>{data.dish[0].name}</h2>
            <h3> ${data.dish[0].price}</h3>
            <h3>Tempo de preparo: {data.dish[0].cookTime} minutos</h3>
            <h3 onClick={ () => {
                if (data.dish[0].category)
                    window.location.pathname = "category/"+data.dish[0].category.id;
            }}>
                Categoria: {data.dish[0].category && data.dish[0].category.name}</h3>
            <h3 onClick={ () => {
                if (data.dish[0].restaurant)
                    window.location.pathname = "restaurant/"+data.dish[0].restaurant.id
            }}>
                Restaurante: {data.dish[0].restaurant && data.dish[0].restaurant.name} - {data.dish[0].restaurant.description}</h3>
        </>
    );
};

export { Dishes, Dish };
