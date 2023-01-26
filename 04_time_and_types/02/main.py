def main(input_value_from_outside=None):
    if input_value_from_outside is None:
        input_value = int(input("Enter a number: "))
    else:
        input_value = input_value_from_outside

    print("Decimal: ", input_value)
    print("Binary: ", bin(input_value))
    print("Hexadecimal: ", hex(input_value))
    return "Decimal: " + str(input_value) + " Binary: " + bin(input_value) + " Hexadecimal: " + hex(input_value) 

if __name__ == "__main__":
    main()
