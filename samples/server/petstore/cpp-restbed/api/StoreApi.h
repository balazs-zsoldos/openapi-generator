/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI-Generator 4.1.3-SNAPSHOT.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/*
 * StoreApi.h
 *
 * 
 */

#ifndef StoreApi_H_
#define StoreApi_H_


#include <memory>
#include <utility>

#include <corvusoft/restbed/session.hpp>
#include <corvusoft/restbed/resource.hpp>
#include <corvusoft/restbed/service.hpp>

#include "Order.h"
#include <map>
#include <string>

namespace org {
namespace openapitools {
namespace server {
namespace api {

using namespace org::openapitools::server::model;

class  StoreApi: public restbed::Service
{
public:
	StoreApi();
	~StoreApi();
	void startService(int const& port);
	void stopService();
};


/// <summary>
/// Delete purchase order by ID
/// </summary>
/// <remarks>
/// For valid response try integer IDs with value &lt; 1000. Anything above 1000 or nonintegers will generate API errors
/// </remarks>
class  StoreApiStoreOrderOrderIdResource: public restbed::Resource
{
public:
	StoreApiStoreOrderOrderIdResource();
    virtual ~StoreApiStoreOrderOrderIdResource();
    void DELETE_method_handler(const std::shared_ptr<restbed::Session> session);
    void GET_method_handler(const std::shared_ptr<restbed::Session> session);

	void set_handler_DELETE(
		std::function<std::pair<int, std::string>(
			std::string const &
		)> handler
	);

	void set_handler_GET(
		std::function<std::pair<int, std::string>(
			int64_t const &
		)> handler
	);

private:
	std::function<std::pair<int, std::string>(
		std::string const &
	)> handler_DELETE_;

	std::function<std::pair<int, std::string>(
		int64_t const &
	)> handler_GET_;
};

/// <summary>
/// Returns pet inventories by status
/// </summary>
/// <remarks>
/// Returns a map of status codes to quantities
/// </remarks>
class  StoreApiStoreInventoryResource: public restbed::Resource
{
public:
	StoreApiStoreInventoryResource();
    virtual ~StoreApiStoreInventoryResource();
    void GET_method_handler(const std::shared_ptr<restbed::Session> session);

	void set_handler_GET(
		std::function<std::pair<int, std::string>(
			
		)> handler
	);


private:
	std::function<std::pair<int, std::string>(
		
	)> handler_GET_;

};

/// <summary>
/// Place an order for a pet
/// </summary>
/// <remarks>
/// 
/// </remarks>
class  StoreApiStoreOrderResource: public restbed::Resource
{
public:
	StoreApiStoreOrderResource();
    virtual ~StoreApiStoreOrderResource();
    void POST_method_handler(const std::shared_ptr<restbed::Session> session);

	void set_handler_POST(
		std::function<std::pair<int, std::string>(
			std::shared_ptr<Order> const &
		)> handler
	);


private:
	std::function<std::pair<int, std::string>(
		std::shared_ptr<Order> const &
	)> handler_POST_;

};


}
}
}
}

#endif /* StoreApi_H_ */

