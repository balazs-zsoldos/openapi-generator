# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six


class TypeHolderDefault(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'string_item': 'str',
        'number_item': 'float',
        'integer_item': 'int',
        'bool_item': 'bool',
        'array_item': 'list[int]',
    }

    attribute_map = {
        'string_item': 'string_item',  # noqa: E501
        'number_item': 'number_item',  # noqa: E501
        'integer_item': 'integer_item',  # noqa: E501
        'bool_item': 'bool_item',  # noqa: E501
        'array_item': 'array_item',  # noqa: E501
    }

    def __init__(self, number_item, integer_item, array_item, string_item='what', bool_item=True):  # noqa: E501
        """TypeHolderDefault - a model defined in OpenAPI

        Args:
            number_item (float):
            integer_item (int):
            array_item (list[int]):

        Keyword Args:
            string_item (str): defaults to 'what', must be one of ['what']  # noqa: E501  # noqa: E501  # noqa: E501
            bool_item (bool): defaults to True, must be one of [True]  # noqa: E501  # noqa: E501
        """

        self._string_item = None
        self._number_item = None
        self._integer_item = None
        self._bool_item = None
        self._array_item = None
        self.discriminator = None

        self.string_item = string_item
        self.number_item = number_item
        self.integer_item = integer_item
        self.bool_item = bool_item
        self.array_item = array_item

    @property
    def string_item(self):
        """Gets the string_item of this TypeHolderDefault.  # noqa: E501


        :return: The string_item of this TypeHolderDefault.  # noqa: E501
        :rtype: str
        """
        return self._string_item

    @string_item.setter
    def string_item(
            self,
            string_item):
        """Sets the string_item of this TypeHolderDefault.


        :param string_item: The string_item of this TypeHolderDefault.  # noqa: E501
        :type: str
        """
        if string_item is None:
            raise ValueError("Invalid value for `string_item`, must not be `None`")  # noqa: E501

        self._string_item = (
            string_item)

    @property
    def number_item(self):
        """Gets the number_item of this TypeHolderDefault.  # noqa: E501


        :return: The number_item of this TypeHolderDefault.  # noqa: E501
        :rtype: float
        """
        return self._number_item

    @number_item.setter
    def number_item(
            self,
            number_item):
        """Sets the number_item of this TypeHolderDefault.


        :param number_item: The number_item of this TypeHolderDefault.  # noqa: E501
        :type: float
        """
        if number_item is None:
            raise ValueError("Invalid value for `number_item`, must not be `None`")  # noqa: E501

        self._number_item = (
            number_item)

    @property
    def integer_item(self):
        """Gets the integer_item of this TypeHolderDefault.  # noqa: E501


        :return: The integer_item of this TypeHolderDefault.  # noqa: E501
        :rtype: int
        """
        return self._integer_item

    @integer_item.setter
    def integer_item(
            self,
            integer_item):
        """Sets the integer_item of this TypeHolderDefault.


        :param integer_item: The integer_item of this TypeHolderDefault.  # noqa: E501
        :type: int
        """
        if integer_item is None:
            raise ValueError("Invalid value for `integer_item`, must not be `None`")  # noqa: E501

        self._integer_item = (
            integer_item)

    @property
    def bool_item(self):
        """Gets the bool_item of this TypeHolderDefault.  # noqa: E501


        :return: The bool_item of this TypeHolderDefault.  # noqa: E501
        :rtype: bool
        """
        return self._bool_item

    @bool_item.setter
    def bool_item(
            self,
            bool_item):
        """Sets the bool_item of this TypeHolderDefault.


        :param bool_item: The bool_item of this TypeHolderDefault.  # noqa: E501
        :type: bool
        """
        if bool_item is None:
            raise ValueError("Invalid value for `bool_item`, must not be `None`")  # noqa: E501

        self._bool_item = (
            bool_item)

    @property
    def array_item(self):
        """Gets the array_item of this TypeHolderDefault.  # noqa: E501


        :return: The array_item of this TypeHolderDefault.  # noqa: E501
        :rtype: list[int]
        """
        return self._array_item

    @array_item.setter
    def array_item(
            self,
            array_item):
        """Sets the array_item of this TypeHolderDefault.


        :param array_item: The array_item of this TypeHolderDefault.  # noqa: E501
        :type: list[int]
        """
        if array_item is None:
            raise ValueError("Invalid value for `array_item`, must not be `None`")  # noqa: E501

        self._array_item = (
            array_item)

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, TypeHolderDefault):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
