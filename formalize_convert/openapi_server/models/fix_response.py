from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from openapi_server.models.base_model import Model
from openapi_server import util


class FixResponse(Model):
    """NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).

    Do not edit the class manually.
    """

    def __init__(self, formal_trs=None):  # noqa: E501
        """FixResponse - a model defined in OpenAPI

        :param formal_trs: The formal_trs of this FixResponse.  # noqa: E501
        :type formal_trs: str
        """
        self.openapi_types = {
            'formal_trs': str
        }

        self.attribute_map = {
            'formal_trs': 'formalTrs'
        }

        self._formal_trs = formal_trs

    @classmethod
    def from_dict(cls, dikt) -> 'FixResponse':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The FixResponse of this FixResponse.  # noqa: E501
        :rtype: FixResponse
        """
        return util.deserialize_model(dikt, cls)

    @property
    def formal_trs(self) -> str:
        """Gets the formal_trs of this FixResponse.


        :return: The formal_trs of this FixResponse.
        :rtype: str
        """
        return self._formal_trs

    @formal_trs.setter
    def formal_trs(self, formal_trs: str):
        """Sets the formal_trs of this FixResponse.


        :param formal_trs: The formal_trs of this FixResponse.
        :type formal_trs: str
        """
        if formal_trs is None:
            raise ValueError("Invalid value for `formal_trs`, must not be `None`")  # noqa: E501

        self._formal_trs = formal_trs
