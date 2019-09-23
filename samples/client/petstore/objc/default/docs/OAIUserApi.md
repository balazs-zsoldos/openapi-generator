# OAIUserApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createUser**](OAIUserApi.md#createuser) | **POST** /user | Create user
[**createUsersWithArrayInput**](OAIUserApi.md#createuserswitharrayinput) | **POST** /user/createWithArray | Creates list of users with given input array
[**createUsersWithListInput**](OAIUserApi.md#createuserswithlistinput) | **POST** /user/createWithList | Creates list of users with given input array
[**deleteUser**](OAIUserApi.md#deleteuser) | **DELETE** /user/{username} | Delete user
[**getUserByName**](OAIUserApi.md#getuserbyname) | **GET** /user/{username} | Get user by user name
[**loginUser**](OAIUserApi.md#loginuser) | **GET** /user/login | Logs user into the system
[**logoutUser**](OAIUserApi.md#logoutuser) | **GET** /user/logout | Logs out current logged in user session
[**updateUser**](OAIUserApi.md#updateuser) | **PUT** /user/{username} | Updated user


# **createUser**
```objc
-(NSURLSessionTask*) createUserWithBody: (OAIUser*) body
        completionHandler: (void (^)(NSError* error)) handler;
```

Create user

This can only be done by the logged in user.

### Example 
```objc

OAIUser* body = [[OAIUser alloc] init]; // Created user object

OAIUserApi*apiInstance = [[OAIUserApi alloc] init];

// Create user
[apiInstance createUserWithBody:body
          completionHandler: ^(NSError* error) {
                        if (error) {
                            NSLog(@"Error calling OAIUserApi->createUser: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OAIUser***](OAIUser.md)| Created user object | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **createUsersWithArrayInput**
```objc
-(NSURLSessionTask*) createUsersWithArrayInputWithBody: (NSArray<OAIUser>*) body
        completionHandler: (void (^)(NSError* error)) handler;
```

Creates list of users with given input array

### Example 
```objc

NSArray<OAIUser>* body = @[[[OAIUser alloc] init]]; // List of user object

OAIUserApi*apiInstance = [[OAIUserApi alloc] init];

// Creates list of users with given input array
[apiInstance createUsersWithArrayInputWithBody:body
          completionHandler: ^(NSError* error) {
                        if (error) {
                            NSLog(@"Error calling OAIUserApi->createUsersWithArrayInput: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NSArray&lt;OAIUser&gt;***](OAIUser.md)| List of user object | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **createUsersWithListInput**
```objc
-(NSURLSessionTask*) createUsersWithListInputWithBody: (NSArray<OAIUser>*) body
        completionHandler: (void (^)(NSError* error)) handler;
```

Creates list of users with given input array

### Example 
```objc

NSArray<OAIUser>* body = @[[[OAIUser alloc] init]]; // List of user object

OAIUserApi*apiInstance = [[OAIUserApi alloc] init];

// Creates list of users with given input array
[apiInstance createUsersWithListInputWithBody:body
          completionHandler: ^(NSError* error) {
                        if (error) {
                            NSLog(@"Error calling OAIUserApi->createUsersWithListInput: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NSArray&lt;OAIUser&gt;***](OAIUser.md)| List of user object | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteUser**
```objc
-(NSURLSessionTask*) deleteUserWithUsername: (NSString*) username
        completionHandler: (void (^)(NSError* error)) handler;
```

Delete user

This can only be done by the logged in user.

### Example 
```objc

NSString* username = @"username_example"; // The name that needs to be deleted

OAIUserApi*apiInstance = [[OAIUserApi alloc] init];

// Delete user
[apiInstance deleteUserWithUsername:username
          completionHandler: ^(NSError* error) {
                        if (error) {
                            NSLog(@"Error calling OAIUserApi->deleteUser: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **NSString***| The name that needs to be deleted | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getUserByName**
```objc
-(NSURLSessionTask*) getUserByNameWithUsername: (NSString*) username
        completionHandler: (void (^)(OAIUser* output, NSError* error)) handler;
```

Get user by user name

### Example 
```objc

NSString* username = @"username_example"; // The name that needs to be fetched. Use user1 for testing.

OAIUserApi*apiInstance = [[OAIUserApi alloc] init];

// Get user by user name
[apiInstance getUserByNameWithUsername:username
          completionHandler: ^(OAIUser* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling OAIUserApi->getUserByName: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **NSString***| The name that needs to be fetched. Use user1 for testing. | 

### Return type

[**OAIUser***](OAIUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **loginUser**
```objc
-(NSURLSessionTask*) loginUserWithUsername: (NSString*) username
    password: (NSString*) password
        completionHandler: (void (^)(NSString* output, NSError* error)) handler;
```

Logs user into the system

### Example 
```objc

NSString* username = @"username_example"; // The user name for login
NSString* password = @"password_example"; // The password for login in clear text

OAIUserApi*apiInstance = [[OAIUserApi alloc] init];

// Logs user into the system
[apiInstance loginUserWithUsername:username
              password:password
          completionHandler: ^(NSString* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling OAIUserApi->loginUser: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **NSString***| The user name for login | 
 **password** | **NSString***| The password for login in clear text | 

### Return type

**NSString***

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **logoutUser**
```objc
-(NSURLSessionTask*) logoutUserWithCompletionHandler: 
        (void (^)(NSError* error)) handler;
```

Logs out current logged in user session

### Example 
```objc


OAIUserApi*apiInstance = [[OAIUserApi alloc] init];

// Logs out current logged in user session
[apiInstance logoutUserWithCompletionHandler: 
          ^(NSError* error) {
                        if (error) {
                            NSLog(@"Error calling OAIUserApi->logoutUser: %@", error);
                        }
                    }];
```

### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateUser**
```objc
-(NSURLSessionTask*) updateUserWithUsername: (NSString*) username
    body: (OAIUser*) body
        completionHandler: (void (^)(NSError* error)) handler;
```

Updated user

This can only be done by the logged in user.

### Example 
```objc

NSString* username = @"username_example"; // name that need to be deleted
OAIUser* body = [[OAIUser alloc] init]; // Updated user object

OAIUserApi*apiInstance = [[OAIUserApi alloc] init];

// Updated user
[apiInstance updateUserWithUsername:username
              body:body
          completionHandler: ^(NSError* error) {
                        if (error) {
                            NSLog(@"Error calling OAIUserApi->updateUser: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **NSString***| name that need to be deleted | 
 **body** | [**OAIUser***](OAIUser.md)| Updated user object | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
