/**
 * 
 * OpenAPI Petstore
 * 
 * 
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 * 
 * Version: 1.0.0
 * 
 * Generated by OpenAPI Generator: https://openapi-generator.tech
 */

// package openapi3graphql-server

// store_api

export default {
    Query: {

        // @return Int!
        GetInventory: () => {
            return {
                
            };
        },

        // @return Order
        GetOrderById: ($orderId) => {
            return {
                "orderId": "789"
            };
        },

    },

    Mutation: {

        // @return 
        DeleteOrder: ($orderId) => {
            return {
                "orderId": "orderId_example"
            };
        },

        // @return Order
        PlaceOrder: ($body) => {
            return {
                "body": ""
            };
        },

    }
}