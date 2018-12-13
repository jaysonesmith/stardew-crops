Feature: Bundles

    Rules:
        - All crops in bundles are returned when no qualifier flag is provided (such as --names, --name, or --crop)
        - Bundle names can be provided to get a list of all crops used in it, if any
        - Crop names can be provided to get a list of all bundles it is used in, if any
        - A list of bundle names can be requested

    Scenario: All crops in bundles are returned when no crop or bundle name is provided
      When bundles is requested without any qualifier flags
      Then all crops used in bundles must be returned
    
    Scenario: A specified bundle that uses crops will return the list of crops it uses

    Scenario: A specified bundle that doesn't use crops will return an informative message


    Scenario: A specified crop used by one or more bundles will return those bundles

    Scenario: A specified crop that isn't used by a bundle will return an informative message

    Scenario Outline: An unmatched bundle or crop will return a not found message

        Examples: 
        |  flag    | value 
        | --bundle | 
        | --crop   |
