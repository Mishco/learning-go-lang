import unittest
from main import main

# python -m unittest .\test_main.py
class TestMain(unittest.TestCase):
    def test_main_five(self):
        # Test input of 5
        input_value = 5
        output = main(input_value)
        self.assertEqual(output, ("Decimal: 5 Binary: 0b101 Hexadecimal: 0x5"))
    
    def test_main_ten(self):
        # Test input of 10
        input_value = 10
        output = main(input_value)
        self.assertEqual(output, ("Decimal: 10 Binary: 0b1010 Hexadecimal: 0xa"))
    
    def test_main_more(self):
        # Test input of 15
        input_value = 15
        output = main(input_value)
        self.assertEqual(output, ("Decimal: 15 Binary: 0b1111 Hexadecimal: 0xf"))

if __name__ == '__main__':
    unittest.main()
