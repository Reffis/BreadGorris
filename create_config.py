import json

if __name__ == "__main__":
    def input_(text):
        return input(f'{text} >> ')

    config_json = dict()

    config_json["token"] = input_("Token")
    config_json["prefix"] = input_("Prefix 1")
    config_json["owner"] = input_("Prefix 2")

    with open(f'config/config.json', 'w', encoding='utf-8') as make_file:
        json.dump(config_json, make_file, indent="\t")