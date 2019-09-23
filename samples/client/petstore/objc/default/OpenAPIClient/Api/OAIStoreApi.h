#import <Foundation/Foundation.h>
#import "OAIOrder.h"
#import "OAIApi.h"

/**
* OpenAPI Petstore
* This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
*
* The version of the OpenAPI document: 1.0.0
* 
*
* NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
* https://openapi-generator.tech
* Do not edit the class manually.
*/



@interface OAIStoreApi: NSObject <OAIApi>

extern NSString* kOAIStoreApiErrorDomain;
extern NSInteger kOAIStoreApiMissingParamErrorCode;

-(instancetype) initWithApiClient:(OAIApiClient *)apiClient NS_DESIGNATED_INITIALIZER;

/// Delete purchase order by ID
/// For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors
///
/// @param orderId ID of the order that needs to be deleted
/// 
///  code:400 message:"Invalid ID supplied",
///  code:404 message:"Order not found"
///
/// @return void
-(NSURLSessionTask*) deleteOrderWithOrderId: (NSString*) orderId
    completionHandler: (void (^)(NSError* error)) handler;


/// Returns pet inventories by status
/// Returns a map of status codes to quantities
///
/// 
///  code:200 message:"successful operation"
///
/// @return NSDictionary<NSString*, NSNumber*>*
-(NSURLSessionTask*) getInventoryWithCompletionHandler: 
    (void (^)(NSDictionary<NSString*, NSNumber*>* output, NSError* error)) handler;


/// Find purchase order by ID
/// For valid response try integer IDs with value <= 5 or > 10. Other values will generated exceptions
///
/// @param orderId ID of pet that needs to be fetched
/// 
///  code:200 message:"successful operation",
///  code:400 message:"Invalid ID supplied",
///  code:404 message:"Order not found"
///
/// @return OAIOrder*
-(NSURLSessionTask*) getOrderByIdWithOrderId: (NSNumber*) orderId
    completionHandler: (void (^)(OAIOrder* output, NSError* error)) handler;


/// Place an order for a pet
/// 
///
/// @param body order placed for purchasing the pet
/// 
///  code:200 message:"successful operation",
///  code:400 message:"Invalid Order"
///
/// @return OAIOrder*
-(NSURLSessionTask*) placeOrderWithBody: (OAIOrder*) body
    completionHandler: (void (^)(OAIOrder* output, NSError* error)) handler;



@end
