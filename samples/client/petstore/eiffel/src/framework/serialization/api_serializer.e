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
	API_SERIALIZER


feature -- Access

	serializer (f: FUNCTION [TUPLE [content_type:STRING; type:ANY],STRING]; a_content_type: STRING; a_type: ANY): STRING
			-- Serialize an object of type `a_type' using the content type `a_content_type'.
		do
			Result := f.item ([a_content_type, a_type])
		end
end
