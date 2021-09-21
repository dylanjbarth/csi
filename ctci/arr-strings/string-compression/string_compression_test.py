from string_compression import compress
import unittest


class TestCompression(unittest.TestCase):

    def test_compression(self):
        """Test that compression code returns expected value"""
        cases = [
            ('a', 'a1'),
            ('aabbccdd', 'a2b2c2d2'),
            ('zaoeu', 'z1a1o1e1u1'),
            ('', ''),
        ]
        for case in cases:
            _input, expected = case
            with self.subTest(case=f"{_input} => {expected}"):
                self.assertEqual(compress(_input), expected)


if __name__ == '__main__':
    unittest.main()
