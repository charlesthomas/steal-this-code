#!/bin/bash

function subcat_menu {
    case $suboption in
    1)
        echo -n kubernetes
        return
        ;;
    2)
        echo -n routing
        return
        ;;
    3)
        echo -n http
        return
        ;;
    4)
        echo -n storage
        return
        ;;
    5)
        echo -n observability
        return
        ;;
    6)
        echo -n cicd
        return
        ;;
    *)
        echo -n $suboption
        return
        ;;
    esac
}

function category_menu {
    case $1 in
    1)
        echo -n "infrastructure-$(subcat_menu $2)"
        return
        ;;
    2)
        echo -n media
        return
        ;;
    3)
        echo -n home-automation
        return
        ;;
    4)
        echo -n tools
        return
        ;;
    5)
        echo -n public
        return
        ;;
    6)
        echo -n misc
        return
        ;;
    *)
        echo -n $1
        return
        ;;
    esac
}

for repo in $(gh repo list --no-archived --source --topic homelab --limit 999 --json name -q .[].name); do
    echo $repo
    echo ---
    echo "1. infrastructure"
    echo "2. media"
    echo "3. home-automation"
    echo "4. tools"
    echo "5. public"
    echo "6. misc"
    echo "or enter manually"
    echo ---
    read option

    if [[ "${option}" == "1" ]]; then
        echo ---
        echo "1. kubernetes"
        echo "2. routing"
        echo "3. http"
        echo "4. storage"
        echo "5. observability"
        echo "6. cicd"
        echo "or enter manually"
        echo ---
        read suboption
    fi

    echo gh repo edit charlesthomas/$repo --add-topic homelab-$(category_menu $option $suboption)
    gh repo edit charlesthomas/$repo --add-topic homelab-$(category_menu $option $suboption)
done
