# coding: utf-8

# flake8: noqa

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

__version__ = "1.0.0"

# import apis into sdk package
from openapi_client.api.another_fake_api import AnotherFakeApi
from openapi_client.api.fake_api import FakeApi
from openapi_client.api.fake_classname_tags_123_api import FakeClassnameTags123Api
from openapi_client.api.pet_api import PetApi
from openapi_client.api.store_api import StoreApi
from openapi_client.api.user_api import UserApi

# import ApiClient
from openapi_client.api_client import ApiClient
from openapi_client.configuration import Configuration
from openapi_client.exceptions import OpenApiException
from openapi_client.exceptions import ApiTypeError
from openapi_client.exceptions import ApiValueError
from openapi_client.exceptions import ApiKeyError
from openapi_client.exceptions import ApiException
# import models into sdk package
from openapi_client.models.additional_properties_any_type import AdditionalPropertiesAnyType
from openapi_client.models.additional_properties_array import AdditionalPropertiesArray
from openapi_client.models.additional_properties_boolean import AdditionalPropertiesBoolean
from openapi_client.models.additional_properties_class import AdditionalPropertiesClass
from openapi_client.models.additional_properties_integer import AdditionalPropertiesInteger
from openapi_client.models.additional_properties_number import AdditionalPropertiesNumber
from openapi_client.models.additional_properties_object import AdditionalPropertiesObject
from openapi_client.models.additional_properties_string import AdditionalPropertiesString
from openapi_client.models.animal import Animal
from openapi_client.models.api_response import ApiResponse
from openapi_client.models.array_of_array_of_number_only import ArrayOfArrayOfNumberOnly
from openapi_client.models.array_of_number_only import ArrayOfNumberOnly
from openapi_client.models.array_test import ArrayTest
from openapi_client.models.capitalization import Capitalization
from openapi_client.models.cat import Cat
from openapi_client.models.cat_all_of import CatAllOf
from openapi_client.models.category import Category
from openapi_client.models.class_model import ClassModel
from openapi_client.models.client import Client
from openapi_client.models.dog import Dog
from openapi_client.models.dog_all_of import DogAllOf
from openapi_client.models.enum_arrays import EnumArrays
from openapi_client.models.enum_class import EnumClass
from openapi_client.models.enum_test import EnumTest
from openapi_client.models.file import File
from openapi_client.models.file_schema_test_class import FileSchemaTestClass
from openapi_client.models.format_test import FormatTest
from openapi_client.models.has_only_read_only import HasOnlyReadOnly
from openapi_client.models.list import List
from openapi_client.models.map_test import MapTest
from openapi_client.models.mixed_properties_and_additional_properties_class import MixedPropertiesAndAdditionalPropertiesClass
from openapi_client.models.model200_response import Model200Response
from openapi_client.models.model_return import ModelReturn
from openapi_client.models.name import Name
from openapi_client.models.number_only import NumberOnly
from openapi_client.models.order import Order
from openapi_client.models.outer_composite import OuterComposite
from openapi_client.models.outer_enum import OuterEnum
from openapi_client.models.pet import Pet
from openapi_client.models.read_only_first import ReadOnlyFirst
from openapi_client.models.special_model_name import SpecialModelName
from openapi_client.models.tag import Tag
from openapi_client.models.type_holder_default import TypeHolderDefault
from openapi_client.models.type_holder_example import TypeHolderExample
from openapi_client.models.user import User
from openapi_client.models.xml_item import XmlItem

