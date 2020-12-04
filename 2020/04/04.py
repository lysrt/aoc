import re

class Passport:
    def __init__(self, raw):
        self.values = {k:v for (k, v) in [x.split(":") for x in raw.split()] }

    def has_required_fields(self):
        required_fields = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
        for f in required_fields:
            if not f in self.values:
                return False
        return True

    def is_valid(self):
        years_ok = validate_year(self.values['byr'], 1920, 2002) and \
            validate_year(self.values['iyr'], 2010, 2020) and \
            validate_year(self.values['eyr'], 2020, 2030)

        height_ok = validate_height(self.values['hgt'])
        hair_ok = validate_color(self.values['hcl'])
        eye_ok = validate_eye_color(self.values['ecl'])
        pid_ok = validate_pid(self.values['pid'])

        return years_ok and height_ok and hair_ok and eye_ok and pid_ok

def validate_pid(pid):
    return bool(re.match('^[0-9]{9}$', pid))

def validate_eye_color(color):
    valid_colors = ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']
    return color in valid_colors

def validate_color(color):
    return bool(re.match('#[0-9a-f]{6}', color))

def validate_height(height):
    if not bool(re.match('[0-9]*(cm|in)', height)):
        return False
    unit = height[-2:] # Keep last 2 chars
    value = int(height[:-2]) # Remove last 2 chars
    return (value >= 150 and value <= 193) if unit == 'cm' else (value >= 59 and value <= 76)

def validate_year(year, low, high):
    if not bool(re.match('[0-9]{4}', year)):
        return False
    return int(year) >= low and int(year) <= high

def main():
    with open('input', 'r') as f:
        raw_passports = f.read()

    passports = [Passport(x) for x in raw_passports.split("\n\n")]

    complete_passports = [p for p in passports if p.has_required_fields()]
    print(len(complete_passports))

    valid_passports = [p for p in complete_passports if p.is_valid()]
    print(len(valid_passports))

main()