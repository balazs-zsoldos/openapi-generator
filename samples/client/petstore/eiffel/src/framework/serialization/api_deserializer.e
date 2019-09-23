note
 description:"[
		OpenAPI Petstore
 		This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
  		The version of the OpenAPI document: 1.0.0
 	    

  	NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).

 		 Do not edit the class manually.
 	]"
	date: "$Date$"
	revision: "$Revision$"
	EIS:"Eiffel openapi generator", "src=https://openapi-generator.tech", "protocol=uri"

class
	API_DESERIALIZER

feature -- Access

	deserializer (f: FUNCTION [TUPLE [content_type:STRING; body:STRING; type:TYPE [detachable ANY]], detachable ANY]; a_content_type: STRING; a_body: STRING; a_type:TYPE [detachable ANY]): detachable ANY
			-- -- From a given response deserialize body `a_body' with conent_type `a_content_type' to a target object of type `a_type'.
		do
			Result := f.item ([a_content_type, a_body, a_type])
		end
end
