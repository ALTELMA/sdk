"""
    Ory Kratos API

    Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.   # noqa: E501

    The version of the OpenAPI document: v0.7.5-alpha.2
    Contact: hi@ory.sh
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import ory_kratos_client
from ory_kratos_client.model.ui_node_anchor_attributes import UiNodeAnchorAttributes
from ory_kratos_client.model.ui_node_image_attributes import UiNodeImageAttributes
from ory_kratos_client.model.ui_node_input_attributes import UiNodeInputAttributes
from ory_kratos_client.model.ui_node_text_attributes import UiNodeTextAttributes
from ory_kratos_client.model.ui_text import UiText
globals()['UiNodeAnchorAttributes'] = UiNodeAnchorAttributes
globals()['UiNodeImageAttributes'] = UiNodeImageAttributes
globals()['UiNodeInputAttributes'] = UiNodeInputAttributes
globals()['UiNodeTextAttributes'] = UiNodeTextAttributes
globals()['UiText'] = UiText
from ory_kratos_client.model.ui_node_attributes import UiNodeAttributes


class TestUiNodeAttributes(unittest.TestCase):
    """UiNodeAttributes unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testUiNodeAttributes(self):
        """Test UiNodeAttributes"""
        # FIXME: construct object with mandatory attributes with example values
        # model = UiNodeAttributes()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
